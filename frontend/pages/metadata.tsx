import Head from 'next/head'

export const metadata = {
    title: 'Taskquire App',
    description: 'Taskquire Description',
}

export default function Metadata() {
    return (
        <Head>
            <title>{metadata.title}</title>
            <meta name="description" content={metadata.description} />
            <link rel="icon" href="/favicon.ico" />
        </Head>
    )
};