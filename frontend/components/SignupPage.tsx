import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import Image from 'next/image';

interface SignupPageProps {}

const SignupPage: React.FC<SignupPageProps> = () => {
    const router = useRouter();

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        const xmls = `<User><Name>${name}</Name><Email>${email}</Email><Password>${password}</Password></User>`

        const response = await fetch('http://localhost:8080/api/users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/xml',
            },
            body: xmls,
        });

        const data = await response.text();

        const parser = new DOMParser();
        const xmlDoc = parser.parseFromString(data,"text/xml");

        const accessToken = xmlDoc.getElementsByTagName("access_token")[0].childNodes[0].nodeValue;
        localStorage.setItem('accessToken', String(accessToken));

        router.push('http://localhost:3000/orgs/your_orgs');
    }

    return (
        <div className="bg-gray-900 min-h-screen flex flex-col items-center justify-center">
            <div className="flex justify-center">
                <Link href="/" passHref>
                    <Image src="/site_logo.png" alt="Taskquire Logo" width={300} height={300} priority />
                </Link>
            </div>
            <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80">
                <h1 className="text-white text-3xl font-bold mb-4">Welcome to Taskquire</h1>
                <p className="text-gray-300 mb-6">Your task management solution</p>
                <form onSubmit={handleSubmit}>
                    <input
                        className="mb-4 w-full px-3 py-2 border border-gray-300 rounded-md" 
                        type="text" 
                        placeholder="Name" 
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                    <input
                        className="mb-4 w-full px-3 py-2 border border-gray-300 rounded-md" 
                        type="email" 
                        placeholder="Email" 
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <input 
                        className="mb-4 w-full px-3 py-2 border border-gray-300 rounded-md" 
                        type="password" 
                        placeholder="Password" 
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <button 
                        className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                        type="submit"
                    >
                        Sign Up
                    </button>
                    <Link href="/login" passHref>
                        <p className="text-blue-500 hover:underline mb-6 cursor-pointer text-center">Click here to login</p>
                    </Link>
                </form>
            </div>
        </div>
    );
};

export default SignupPage;