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
    OwnerName: string[];
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

interface OneTeamAllProjectsPageProps {}

const OneTeamAllProjectsPage: React.FC<OneTeamAllProjectsPageProps> = () => {
    const [team, setTeam] = useState<Team | null>(null);
    const [projects, setProjects] = useState<Project[]>([]);
    const router = useRouter();

    const { teamName } = router.query;

    useEffect(() => {
        const fetchTeam = async () => {
            const response = await fetch(`http://localhost:8080/api/teams/${teamName}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();

            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setTeam(result.teams.team);
                }
            });
        }

        if (teamName) {
            fetchTeam();
        }
    }, [teamName]);

    useEffect(() => {
        const fetchProjects = async () => {
            if (!teamName) {
                return;
            }

            const response = await fetch(`http://localhost:8080/api/teams/${teamName}/projects`, {
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
    }, [teamName]);

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
            {team && (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4">
                    <h1 className="text-white text-3xl font-bold mb-4">{team.Name}</h1>
                    <p className="text-gray-300 mb-6"><strong>Description:</strong> {team.Description.String}</p>
                    <p className="text-gray-300 mb-6"><strong>Owner:</strong> {team.OwnerName}</p>
                    <p className="text-gray-300 mb-6"><strong>Created At:</strong> {team.CreatedAt}</p>
                    <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {team.UpdatedAt}</p>
                </div>
            )}
            {team && projects && (Array.isArray(projects) ? projects : [projects]).map((project, index) => (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4" key={index}>
                    <div className="text-white flex justify-between">
                        <Link href={{ pathname: "/orgs/[orgName]/projects/[projectName]/update",
                                         query: { orgName: project.OrgName, projectName: project.Name, projectDescription: project.Description } }}>
                            <button>✏️</button>
                        </Link>
                        <button onClick={() => deleteProject(String(project.OrgName), String(project.Name))}>X</button>
                    </div>
                    <Link href="/orgs/[orgName]/projects/[projectName]/teams/[teamName]/tasks" as={`/orgs/${project.OrgName}/projects/${project.Name}/teams/${team.Name}/tasks`} key={index}>
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

export default OneTeamAllProjectsPage;