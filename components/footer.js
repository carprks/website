import Link from 'next/link'

const footerStyle = {
    width: '33%',
    float: 'left'
};
const linkStyle = {
    float: 'left',
    clear: 'both',
    width: '100%'
};

export default function Footer() {
    return (
        <footer>
            <div style={footerStyle}>
                <Link href="/about">
                    <a style={linkStyle}>About</a>
                </Link>
                <Link href="/privacy">
                    <a style={linkStyle}>Privacy</a>
                </Link>
                <Link href="/contact">
                    <a style={linkStyle}>Contact Us</a>
                </Link>
            </div>
        </footer>
    )
}
