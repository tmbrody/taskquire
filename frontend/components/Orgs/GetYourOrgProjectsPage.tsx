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

interface Project {
    Name: string[];
    Description: string[];
    OrgName: string[];
    CreatedAt: string[];
    UpdatedAt: string[];
}

interface OneOrgAllProjectsPageProps {}

const OneOrgAllProjectsPage: React.FC<OneOrgAllProjectsPageProps> = () => {
    const [org, setOrg] = useState<Org | null>(null);
    const [projects, setProjects] = useState<Project[]>([]);
    const router = useRouter();

    const { orgName } = router.query;

    useEffect(() => {
        const fetchOrg = async () => {
            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();

            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setOrg({
                        Name: result.Org.Name,
                        Description: result.Org.Description,
                        OwnerID: result.Org.OwnerID,
                        CreatedAt: result.Org.CreatedAt,
                        UpdatedAt: result.Org.UpdatedAt,
                    });
                }
            });
        }

        if (orgName) {
            fetchOrg();
        }
    }, [orgName]);

    useEffect(() => {
        const fetchProjects = async () => {
            if (!orgName) {
                return;
            }

            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();
            
            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setProjects(result.projects.project);
                }
            });
        }

        fetchProjects();
    }, [orgName]);

    const deleteProject = async (orgName: string, projectName: string) => {
        const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/xml',
                'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
        });

        if (response.ok) {
            router.reload();
        } else {
            console.error('Failed to delete project');
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
            {org && (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4">
                    <h1 className="text-white text-3xl font-bold mb-4">{org.Name}</h1>
                    <p className="text-gray-300 mb-6"><strong>Description:</strong> {org.Description}</p>
                    <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {org.OwnerID}</p>
                    <p className="text-gray-300 mb-6"><strong>Created At:</strong> {org.CreatedAt}</p>
                    <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {org.UpdatedAt}</p>
                </div>
            )}
            {org && projects && (Array.isArray(projects) ? projects : [projects]).map((project, index) => (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4" key={index}>
                    <div className="text-white flex justify-between">
                        <Link href={{ pathname: "/orgs/[orgName]/projects/[projectName]/update",
                                         query: { orgName: orgName, projectName: project.Name, projectDescription: project.Description } }}>
                            <button>✏️</button>
                        </Link>
                        <button onClick={() => deleteProject(String(org.Name), String(project.Name))}>X</button>
                    </div>
                    <Link href="/orgs/[orgName]/projects/[projectName]/teams" as={`/orgs/${project.OrgName}/projects/${project.Name}/teams`} key={index}>
                        <h1 className="text-white text-3xl font-bold mb-4">{project.Name}</h1>
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {project.Description}</p>
                        <p className="text-gray-300 mb-6"><strong>Organization:</strong> {project.OrgName}</p>
                        <p className="text-gray-300 mb-6"><strong>Created At:</strong> {project.CreatedAt}</p>
                        <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {project.UpdatedAt}</p>
                    </Link>
                </div>
            ))}
        </div>
    );
};

export default OneOrgAllProjectsPage;