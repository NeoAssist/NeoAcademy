import React from 'react';
import { element } from 'prop-types';
import { Container } from './styles';
import Nav from '../../components/Nav';

const AuthLayout = ({ children }) => (
    <Container>
        <Nav />
        {children}
    </Container>
);

AuthLayout.propTypes = {
    children: element.isRequired,
};

export default AuthLayout;
