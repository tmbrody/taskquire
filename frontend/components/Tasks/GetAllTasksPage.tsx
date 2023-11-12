import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import Image from 'next/image';
import xml2js from 'xml2js';
import { useRouter } from 'next/router';

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
    Users: string[];
}

interface Task {
    Name: string[];
    Description: Description;
    OwnerID: string[];
    ProjectID: string[];
    TeamID: string[];
    ParentID: string[];
    CreatedAt: string[];
    UpdatedAt: string[];
}

const OneProjectPage: React.FC = () => {
    const [team, setTeam] = useState<Team | null>(null);
    const [tasks, setTasks] = useState<Task | null>(null);

    const router = useRouter();

    const { orgName, projectName, teamName } = router.query;

    useEffect(() => {
        const fetchTeam = async () => {
            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}/teams/${teamName}`, {
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
        const fetchTasks = async () => {
            if (!teamName) {
                return;
            }

            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}/teams/${teamName}/tasks`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();
            
            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setTasks(result.tasks.task);
                }
            });
        }

        fetchTasks();
    }, [teamName]);

    const deleteTask = async (taskName: string) => {

        const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}/teams/${teamName}/tasks/${taskName}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/xml',
                'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
            }
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
            {team && (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4">
                    <h1 className="text-white text-3xl font-bold mb-4">{team.Name}</h1>
                    {String(team.Description.Valid) === 'true' ? (
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {team.Description.String}</p>) : (
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {null}</p>)
                    }
                    <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {team.OwnerID}</p>
                    <p className="text-gray-300 mb-6"><strong>Created At:</strong> {team.CreatedAt}</p>
                    <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {team.UpdatedAt}</p>
                    <p className="text-gray-300 mb-6"><strong>Users:</strong> {team.Users}</p>
                </div>
            )}
            {team && tasks && (Array.isArray(tasks) ? tasks : [tasks]).map((task, index) => (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4" key={index}>
                    <div className="text-white flex justify-between">
                        <Link href={{ pathname: "/orgs/[orgName]/projects/[projectName]/teams/[teamName]/tasks/[taskName]/update", 
                                        query: { orgName: orgName, projectName: projectName, teamName: teamName, taskName: task.Name, taskDescription: task.Description.String } }}>
                            <button>✏️</button>
                        </Link>
                        <button onClick={() => deleteTask(String(task.Name))}>X</button>
                    </div>
                    <h1 className="text-white text-3xl font-bold mb-4">{task.Name}</h1>
                    {String(task.Description.Valid) === 'true' ? (
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {task.Description.String}</p>) : (
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {null}</p>)
                    }
                    <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {task.OwnerID}</p>
                    <p className="text-gray-300 mb-6"><strong>Created At:</strong> {task.CreatedAt}</p>
                    <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {task.UpdatedAt}</p>
                </div>
            ))}
            {team &&
                <Link href="/orgs/[orgName]/projects/[projectName]/teams/[teamName]/tasks/create" 
                    as={`/orgs/${orgName}/projects/${projectName}/teams/${team.Name}/tasks/create`} passHref>
                    <button 
                        className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                        type="submit"
                    >
                        Create Task
                    </button>
                </Link>
            }
            {team &&
                <Link href="/orgs/[orgName]/projects/[projectName]/teams/[teamName]/adduser"
                    as={`/orgs/${orgName}/projects/${projectName}/teams/${team.Name}/adduser`} passHref>
                    <button 
                        className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                        type="submit"
                    >
                        Add User to Team
                    </button>
                </Link>
            }
            {team &&
                <Link href="/orgs/[orgName]/projects/[projectName]/teams/[teamName]/removeuser"
                    as={`/orgs/${orgName}/projects/${projectName}/teams/${team.Name}/removeuser`} passHref>
                    <button 
                        className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                        type="submit"
                    >
                        Remove User to Team
                    </button>
                </Link>
            }
        </div>
    );
};

export default OneProjectPage;