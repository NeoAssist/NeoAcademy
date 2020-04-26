import React from 'react';
import { string, bool, element, oneOfType, shape } from 'prop-types';
import AuthLayout from '../layouts/auth';
import DefaultLayout from '../layouts/default';
import Head from '../components/head';
import GlobalStyle from '../styles/global';

const Layout = ({ signed, children }) => {
    const Content = signed ? AuthLayout : DefaultLayout;

    return <Content>{children}</Content>;
};

const App = ({ Component, pageProps, router }) => {
    const signed = true;
    if (process.browser) {
        switch (router.pathname) {
            case '/':
                router.push(signed ? '/dashboard' : '/login');
                break;
            case '/login':
                if (signed) {
                    router.push('/dashboard');
                }
                break;
            default:
                if (!signed) {
                    router.push('/login');
                }
                break;
        }
    }

    return (
        <Layout signed={signed}>
            <Head />
            <>
                <Component {...pageProps} />
                <GlobalStyle />
            </>
        </Layout>
    );
};

Layout.propTypes = {
    signed: bool.isRequired,
    children: oneOfType([string, element]).isRequired,
};

App.propTypes = {
    Component: element.isRequired,
    pageProps: shape.isRequired,
    router: shape.isRequired,
};

export default App;
