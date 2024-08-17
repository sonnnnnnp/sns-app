import Link from 'next/link';

export default function Navigation() {
    const navItems = [
        { href: '/home', label: 'Home' },
        { href: '/chat', label: 'Chat' },
        { href: '/notifications', label: 'Notifications' },
        { href: '/profile', label: 'Profile'},
        { href: '/action', label: 'Action'},
    ];
    return (
        <nav>
            <ul>
            {navItems.map( item => (
                <li>
                    <Link href={item.href}>
                        <a>{item.label}</a>
                    </Link>
                </li>
            ))}                
            </ul>
        </nav>
    )
}