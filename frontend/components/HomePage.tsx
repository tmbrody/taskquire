import Link from 'next/link';
import Image from 'next/image';

interface HomePageProps {}

const HomePage: React.FC<HomePageProps> = () => {
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
            <div className="bg-gray-800 p-8 rounded-lg shadow-md">
                <h1 className="text-white text-3xl font-bold mb-4">Welcome to Taskquire</h1>
                <p className="text-gray-300 mb-6">Your task management solution</p>
                <Link href="/orgs/all_orgs" passHref>
                    <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg">
                        View All Orgs
                    </button>
                </Link>
                <Link href="/orgs/your_orgs" passHref>
                    <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg">
                        View Your Orgs
                    </button>
                </Link>
                <Link href="/teams/all_teams" passHref>
                    <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg">
                        View All Teams
                    </button>
                </Link>
                <Link href="/teams/your_teams" passHref>
                    <button className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-full hover:shadow-lg">
                        View Your Teams
                    </button>
                </Link>
            </div>
        </div>
    );
};

export default HomePage;
