import { createGlobalStyle } from 'styled-components';

export default createGlobalStyle`
    * {
        margin: 0;
        padding: 0;
        outline: 0;
        box-sizing: border-box;
    }

    *:focus {
        outline: 0;
    }

    html,
    body,
    #__next {
        height: 100%;
    }

    body {
        -webkit-font-smoothing: antialiased;
    }

    body,
    input,
    button {
        font: 14px 'Open Sans', sans-serif;
    }

    a {
        text-decoration: none;
    }

    ul {
        list-style: none;
    }

    button {
        background: #fff;
        border: 0;
        cursor: pointer;
        outline: 0;
    }
`;
