import Header from './header'
import Footer from './footer'

const layoutStyle = {
    margin: 20,
    padding: 20,
    border: '1px solid #DDD'
}

const Layout = Page => {
    return () => (
        <div style={layoutStyle}>
            <Header/>
            <Page/>
            <Footer/>
        </div>
    )
}

export default Layout
