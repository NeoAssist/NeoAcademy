import React from 'react';
import { string, element, oneOfType } from 'prop-types';
import { Container } from './styles';
import Nav from '../../components/Nav';

const AuthLayout = ({ children }) => (
    <Container>
        <Nav />
        {children}
    </Container>
);

AuthLayout.propTypes = {
    children: oneOfType([string, element]).isRequired,
};

export default AuthLayout;
