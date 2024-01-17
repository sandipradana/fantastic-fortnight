import {  ScrollArea } from '@mantine/core';
import {
    IconDashboard,
    IconBurger,
    IconUser,
} from '@tabler/icons-react';

import classes from './navbar.module.css';
import { LinkGroup } from './linkgroup';

const mockdata = [
    { label: 'Dashboard', icon: IconDashboard, href: '/' },
    { label: 'Products', icon: IconBurger, href: '/products' },
    { label: 'Admins', icon: IconUser, href: '/admins' },
];

export function Navbar() {
    const links = mockdata.map((item) => <LinkGroup {...item} key={item.label} />);

    return (
        <nav className={classes.navbar}>
            <ScrollArea className={classes.links}>
                <div className={classes.linksInner}>{links}</div>
            </ScrollArea>
        </nav>
    );
}