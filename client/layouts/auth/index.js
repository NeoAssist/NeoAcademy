import React from 'react';
import { Container } from './styles';
import Nav from '../../components/Nav';

const AuthLayout = ({ children }) => (
    <Container>
        <Nav />
        {children}
    </Container>
);

export default AuthLayout;
