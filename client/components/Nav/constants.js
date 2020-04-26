export const navLinks = [
    { href: '/dashboard', label: 'Dashboard' },
    { href: '/profile', label: 'Profile' },
    { href: '/login', label: 'Login' },
].map((link) => {
    link.key = `nav-link-${link.href}-${link.label}`;
    return link;
});
