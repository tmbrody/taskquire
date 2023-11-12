import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import Image from 'next/image';
import xml2js from 'xml2js';

interface Org {
    Name: string[];
    Description: string[];
    OwnerID: string[];
    CreatedAt: string[];
    UpdatedAt: string[];
}

interface YourOrgsPageProps {}

const YourOrgsPage: React.FC<YourOrgsPageProps> = () => {
    const [orgs, setOrgs] = useState<Org[]>([]);

    const router = useRouter();

    useEffect(() => {
        const fetchOrgs = async () => {
            const response = await fetch('http://localhost:8080/api/orgs', {
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
                        setOrgs(result.orgs.org);
                    }
                });
            }
        }

        fetchOrgs();
    }, []);

    const deleteOrg = async (orgName: string) => {
        const response = await fetch(`http://localhost:8080/api/orgs/${orgName}`, {
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
            {orgs && (Array.isArray(orgs) ? orgs : [orgs]).map((org, index) => (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4" key={index}>
                    <div className="text-white flex justify-between">
                        <Link href={{ pathname: "/orgs/[orgName]/update", 
                                        query: { orgName: org.Name, orgDescription: org.Description } }}>
                            <button>✏️</button>
                        </Link>
                        <button onClick={() => deleteOrg(String(org.Name))}>X</button>
                    </div>
                    <Link href="/orgs/[orgName]/projects" as={`/orgs/${org.Name}/projects`}>
                        <h1 className="text-white text-3xl font-bold mb-4">{org.Name}</h1>
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {org.Description}</p>
                        <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {org.OwnerID}</p>
                        <p className="text-gray-300 mb-6"><strong>Created At:</strong> {org.CreatedAt}</p>
                        <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {org.UpdatedAt}</p>
                    </Link>
                </div>
            ))}
            <Link href="/orgs/create" passHref>
                <button 
                    className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                    type="submit"
                >
                    Create Org
                </button>
            </Link>
        </div>
    );
};

export default YourOrgsPage;