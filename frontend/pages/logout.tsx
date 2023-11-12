import { useRouter } from 'next/router';

export default function Logout() {
    const router = useRouter();

    localStorage.setItem('accessToken', "");

    router.push('http://localhost:3000/');
};