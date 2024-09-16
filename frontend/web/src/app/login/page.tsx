import { Container } from "@/components/section/container";

export default function Home(){
    return (
        <div>
            <Container>
                <div className="w-full h-screen justify-center bg-primary-blue bg-opacity-50">
                    <span className="text-blue-900 font-bold text-3xl bg-white bg-center border border-blue-900 justify-center">Login</span>
                </div>
            </Container>
        </div>
    )
}