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
        <div className="w-52">
            <div className="fixed">
                <nav className="text-lg">
                    <ul>
                    {navItems.map( item => (
                        <li className="w-52 p-3">
                            <Link href={item.href}>{item.label}</Link>
                        </li>
                    ))}                
                    </ul>
                </nav>
            </div>
        </div>
    )
}