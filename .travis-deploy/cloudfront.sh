#!/usr/bin/env bash
DEPLOY_ENV=dev

invalid()
{
    id=""

    dists=$(AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY aws cloudfront list-distributions)
    items=$(echo $dists | jq -c .DistributionList.Items[])
    for item in $items; do
        website=$(echo $item | jq -c .Origins.Items[].Id)
        website=$(echo $website | sed -e 's/^"//' -e 's/"$//')
        if [[ $website == "website-cloudfront-$DEPLOY_ENV" ]]; then
            id=$(echo $item | jq .Id)
        fi
    done

    if [[ ! -z $id ]]; then
        id=$(echo $id | sed -e 's/^"//' -e 's/"$//')
        AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY aws cloudfront create-invalidation --paths "/*" --distribution-id $id
    fi
}


if [[ -z "$TRAVIS_PULL_REQUEST" ]] || [[ "$TRAVIS_PULL_REQUEST" == "false" ]]; then
    AWS_ACCESS_KEY_ID=$DEV_AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY=$DEV_AWS_SECRET_ACCESS_KEY
    AWS_CLOUDFRONT_ID=$DEV_AWS_CLOUDFRONT_ID

    echo "Invalidating Dev CloudFront"
    invalid
    echo "Invalidated Dev CloudFront"

    if [[ -z "$SKIP_LIVE" ]] || [[ "$SKIP_LIVE" == "false" ]]; then
        if [[ "$TRAVIS_BRANCH" == "master" ]]; then
            AWS_ACCESS_KEY_ID=$LIVE_AWS_ACCESS_KEY_ID
            AWS_SECRET_ACCESS_KEY=$LIVE_AWS_SECRET_ACCESS_KEY
            AWS_CLOUDFRONT_ID=$LIVE_AWS_CLOUDFRONT_ID

            echo "Invalidating Live CloudFront"
            invalid
            echo "Invalidated Live CloudFront"
        fi
    fi
fi