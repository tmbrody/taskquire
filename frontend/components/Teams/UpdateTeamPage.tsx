import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import Image from 'next/image';

interface UpdateTeamPageProps {}

const UpdateTeamPage: React.FC<UpdateTeamPageProps> = () => {
    const router = useRouter();

    const { teamName } = router.query;
    const { teamDescription } = router.query;

    const [name, setName] = useState(teamName || '');
    const [description, setDescription] = useState(teamDescription || '');

    const handleUpdateTeam = async (event: React.FormEvent) => {
        event.preventDefault();

        const xmls = `<Team><Name>${name}</Name><Description>${description}</Description></Team>`

        const response = await fetch(`http://localhost:8080/api/teams/${teamName}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/xml',
                'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: xmls,
        });

        router.push(`http://localhost:3000/teams/your_teams`)
    }

    return (
        <div className="bg-gray-900 min-h-screen flex flex-col items-center justify-center">
            <div className="flex flex-col items-center justify-center">
                <Link href="/" passHref>
                    <Image src="/site_logo.png" alt="Taskquire Logo" width={300} height={300} priority />
                </Link>
                <div className="flex flex-row items-center justify-center">
                    <Link href="/signup" passHref>
                        <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg">
                            Sign Up
                        </button>
                    </Link>
                    <Link href="/login" passHref>
                        <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg">
                            Login
                        </button>
                    </Link>
                    <Link href="/logout" passHref>
                        <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg">
                            Logout
                        </button>
                    </Link>
                </div>
            </div>
            <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80">
                <h1 className="text-white text-3xl font-bold mb-4">Welcome to Taskquire</h1>
                <p className="text-gray-300 mb-6">Your task management solution</p>
                <form onSubmit={handleUpdateTeam}>
                    <input
                        className="mb-4 w-full px-3 py-2 border border-gray-300 rounded-md" 
                        type="text" 
                        placeholder="Name" 
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                    <input 
                        className="mb-4 w-full px-3 py-2 border border-gray-300 rounded-md" 
                        type="text"
                        placeholder="Description" 
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                    <button 
                        className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                        type="submit"
                    >
                        Update Team
                    </button>
                </form>
            </div>
        </div>
    );
};

export default UpdateTeamPage;