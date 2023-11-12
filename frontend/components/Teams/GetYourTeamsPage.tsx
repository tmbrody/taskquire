import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import Image from 'next/image';
import xml2js from 'xml2js';

interface Description {
    String: string[];
    Valid: string[];
}

interface Team {
    Name: string[];
    Description: Description;
    OwnerID: string[];
    CreatedAt: string[];
    UpdatedAt: string[];
}

interface YourTeamsPageProps {}

const YourTeamsPage: React.FC<YourTeamsPageProps> = () => {
    const [teams, setTeams] = useState<Team[]>([]);

    const router = useRouter();

    useEffect(() => {
        const fetchTeams = async () => {
            const response = await fetch('http://localhost:8080/api/teams', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            if (localStorage.getItem('accessToken') === '' || response.status === 401) {
                router.push('http://localhost:3000/login');
            } else {
                const data = await response.text();

                xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                    if (!err) {
                        setTeams(result.teams.team);
                    }
                });
            }
        }

        fetchTeams();
    }, []);

    const deleteOrg = async (teamName: string) => {
        const response = await fetch(`http://localhost:8080/api/teams/${teamName}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/xml',
                'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
        });

        if (response.ok) {
            router.reload();
        } else {
            console.error('Failed to delete org');
        }
    };

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
            {teams && (Array.isArray(teams) ? teams : [teams]).map((team, index) => (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4" key={index}>
                    <div className="text-white flex justify-between">
                        <Link href={{ pathname: "/teams/[teamName]/update", 
                                        query: { teamName: team.Name, teamDescription: team.Description.String } }}>
                            <button>✏️</button>
                        </Link>
                        <button onClick={() => deleteOrg(String(team.Name))}>X</button>
                    </div>
                    <Link href="/teams/[teamName]/projects" as={`/teams/${team.Name}/projects`} key={index}>
                        <h1 className="text-white text-3xl font-bold mb-4">{team.Name}</h1>
                        {String(team.Description.Valid) === 'true' ? (
                            <p className="text-gray-300 mb-6"><strong>Description:</strong> {team.Description.String}</p>) : (
                            <p className="text-gray-300 mb-6"><strong>Description:</strong> {null}</p>)
                        }
                        <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {team.OwnerID}</p>
                        <p className="text-gray-300 mb-6"><strong>Created At:</strong> {team.CreatedAt}</p>
                        <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {team.UpdatedAt}</p>
                    </Link>
                </div>
            ))}
            <Link href="/teams/create" passHref>
                <button 
                    className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                    type="submit"
                >
                    Create Team
                </button>
            </Link>
        </div>
    );
};

export default YourTeamsPage;