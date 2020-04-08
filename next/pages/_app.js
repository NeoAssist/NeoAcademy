import Head from '../../components/head';

function MyApp({ Component, pageProps, router }) {
  console.log('teste')
  console.log(router)
  return (
    <>
      <Head />
      <Component {...pageProps} />
    </>
  )

}

export default MyApp;
