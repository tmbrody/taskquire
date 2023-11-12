import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import Image from 'next/image';
import xml2js from 'xml2js';
import { useRouter } from 'next/router';

interface Description {
    String: string[];
    Valid: string[];
}

interface Project {
    Name: string[];
    Description: string[];
    OrgID: string[];
    CreatedAt: string[];
    UpdatedAt: string[];
}

interface Team {
    Name: string[];
    Description: Description;
    OwnerID: string[];
    CreatedAt: string[];
    UpdatedAt: string[];
}

const OneProjectPage: React.FC = () => {
    const [project, setProject] = useState<Project | null>(null);
    const [teams, setTeams] = useState<Team | null>(null);

    const router = useRouter();

    const { orgName, projectName } = router.query;

    useEffect(() => {
        const fetchProject = async () => {
            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();

            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setProject(result.projects.project);
                }
            });
        }

        if (projectName) {
            fetchProject();
        }
    }, [projectName]);

    useEffect(() => {
        const fetchTeams = async () => {
            if (!projectName) {
                return;
            }

            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}/teams`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();
            
            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setTeams(result.teams.team);
                }
            });
        }

        fetchTeams();
    }, [projectName]);

    const removeTeamFromProject = async (projectName: string, teamName: string) => {
        const xmls = `<Combine><Project>${projectName}</Project><Team>${teamName}</Team></Combine>`

        const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}/teams`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/xml',
                'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            },
            body: xmls
        });

        if (response.ok) {
            router.reload();
        } else {
            console.error('Failed to remove team from project');
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
            {project && (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4">
                    <h1 className="text-white text-3xl font-bold mb-4">{project.Name}</h1>
                    <p className="text-gray-300 mb-6"><strong>Description:</strong> {project.Description}</p>
                    <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {project.OrgID}</p>
                    <p className="text-gray-300 mb-6"><strong>Created At:</strong> {project.CreatedAt}</p>
                    <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {project.UpdatedAt}</p>
                </div>
            )}
            {project && teams && (Array.isArray(teams) ? teams : [teams]).map((team, index) => (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4" key={index}>
                    <div className="text-white text-right">
                        <button onClick={() => removeTeamFromProject(String(project.Name), String(team.Name))}>X</button>
                    </div>
                    <Link href="/orgs/[orgName]/projects/[projectName]/teams/[teamName]/tasks"
                         as={`/orgs/${orgName}/projects/${project.Name}/teams/${team.Name}/tasks`} key={index}>
                        <h1 className="text-white text-3xl font-bold mb-4">{team.Name}</h1>
                        {String(team.Description.Valid) === 'true' ? (
                            <p className="text-gray-300 mb-6"><strong>Description:</strong> {team.Description.String}</p>) : (
                            <p className="text-gray-300 mb-6"><strong>Description:</strong> {null}</p>)
                        }
                        <p className="text-gray-300 mb-6"><strong>Org ID:</strong> {team.OwnerID}</p>
                        <p className="text-gray-300 mb-6"><strong>Created At:</strong> {team.CreatedAt}</p>
                        <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {team.UpdatedAt}</p>
                    </Link>
                </div>
            ))}
            {project &&
                <Link href="/orgs/[orgName]/projects/[projectName]/teams/add" as={`/orgs/${orgName}/projects/${project.Name}/teams/addtoproject`} passHref>
                    <button 
                        className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                        type="submit"
                    >
                        Add Team to Project
                    </button>
                </Link>
            }
        </div>
    );
};

export default OneProjectPage;