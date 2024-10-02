import { Header_home } from "../../components/header/header_home/header_home";

export default function Layout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <>
            <Header_home />
            {children}
        </>
    );
}
