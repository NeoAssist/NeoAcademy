import React, { Fragment } from 'react';
import AuthLayout from '../layouts/auth';
import DefaultLayout from '../layouts/default';
import Head from '../components/head';
import GlobalStyle from '../styles/global';

const Layout = ({ signed, children }) => {
    let Content = signed ? AuthLayout : DefaultLayout;

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

export default App;
