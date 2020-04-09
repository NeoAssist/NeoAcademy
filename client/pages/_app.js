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
    const signed = false;
    if (process.browser) {
        do {
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
        } while (0);
    }

    return (
        <>
            <Head />
            <Layout signed={signed}>
                <Component {...pageProps} />
                <GlobalStyle />
            </Layout>
        </>
    );
};

export default App;