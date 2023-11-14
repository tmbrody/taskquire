import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import Image from 'next/image';
import xml2js from 'xml2js';
import { useRouter } from 'next/router';

interface Description {
    String: string[];
    Valid: string[];
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
    Subtasks: string[];
}

const OneTaskPage: React.FC = () => {
    const [task, setTask] = useState<Task | null>(null);
    const [subtasks, setSubtasks] = useState<Task | null>(null);

    const router = useRouter();

    const { orgName, projectName, teamName, taskName } = router.query;

    useEffect(() => {
        const fetchTask = async () => {
            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}/teams/${teamName}/tasks/${taskName}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();

            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setTask(result.tasks.task);
                }
            });
        }

        if (taskName) {
            fetchTask();
        }
    }, [taskName]);

    useEffect(() => {
        const fetchSubtasks = async () => {
            if (!taskName) {
                return;
            }

            const response = await fetch(`http://localhost:8080/api/orgs/${orgName}/projects/${projectName}/teams/${teamName}/tasks/${taskName}/subtasks`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/xml',
                    'Authorization': `Bearer ${localStorage.getItem('accessToken')}`
                },
            });

            const data = await response.text();
            console.log(data);
            
            xml2js.parseString(data, { explicitArray: false }, (err, result) => {
                if (!err) {
                    setSubtasks(result.tasks.task);
                }
            });
        }

        fetchSubtasks();
    }, [taskName]);

    const deleteSubtask = async (taskName: string) => {

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
            {task && (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4">
                    <h1 className="text-white text-3xl font-bold mb-4">{task.Name}</h1>
                    {String(task.Description.Valid) === 'true' ? (
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {task.Description.String}</p>) : (
                        <p className="text-gray-300 mb-6"><strong>Description:</strong> {null}</p>)
                    }
                    <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {task.OwnerID}</p>
                    <p className="text-gray-300 mb-6"><strong>Created At:</strong> {task.CreatedAt}</p>
                    <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {task.UpdatedAt}</p>
                </div>
            )}
            {task && subtasks && (Array.isArray(subtasks) ? subtasks : [subtasks]).map((subtask, index) => (
                <div className="bg-gray-800 p-8 rounded-lg shadow-md w-80 mb-4" key={index}>
                    <div className="text-white flex justify-between">
                        <Link href={{ pathname: "/orgs/[orgName]/projects/[projectName]/teams/[teamName]/tasks/[taskName]/subtasks/[subtaskName]/update", 
                                        query: { orgName: orgName, projectName: projectName, teamName: teamName, taskName: task.Name, 
                                                subtaskName: subtask.Name, subtaskDescription: subtask.Description.String } }}>
                            <button>✏️</button>
                        </Link>
                        <button onClick={() => deleteSubtask(String(subtask.Name))}>X</button>
                    </div>
                    <Link href="/orgs/[orgName]/projects/[projectName]/teams/[teamName]/tasks/[taskName]/subtasks"
                            as={`/orgs/${orgName}/projects/${projectName}/teams/${teamName}/tasks/${subtask.Name}/subtasks`} key={index}>
                        <h1 className="text-white text-3xl font-bold mb-4">{subtask.Name}</h1>
                        {String(subtask.Description.Valid) === 'true' ? (
                            <p className="text-gray-300 mb-6"><strong>Description:</strong> {subtask.Description.String}</p>) : (
                            <p className="text-gray-300 mb-6"><strong>Description:</strong> {null}</p>)
                        }
                        <p className="text-gray-300 mb-6"><strong>Owner ID:</strong> {subtask.OwnerID}</p>
                        <p className="text-gray-300 mb-6"><strong>Created At:</strong> {subtask.CreatedAt}</p>
                        <p className="text-gray-300 mb-6"><strong>Updated At:</strong> {subtask.UpdatedAt}</p>
                        <p className="text-gray-300 mb-6"><strong>Subtasks:</strong> {subtask.Subtasks}</p>
                    </Link>
                </div>
            ))}
            {task &&
                <Link href="/orgs/[orgName]/projects/[projectName]/teams/[teamName]/tasks/[taskName]/subtasks/create" 
                    as={`/orgs/${orgName}/projects/${projectName}/teams/${teamName}/tasks/${taskName}/subtasks/create`} passHref>
                    <button 
                        className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg w-full" 
                        type="submit"
                    >
                        Create Subtask
                    </button>
                </Link>
            }
        </div>
    );
};

export default OneTaskPage;