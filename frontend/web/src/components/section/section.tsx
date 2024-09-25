import "next/image"
import { Container } from "./container"
export function Section(){
    return (
        <section className="w-full h-screen bg-background-estacionamento bg-no-repeat bg-center bg-cover">
            <Container>
                <div className="container bg-primary-blue h-80 w-[500px] bg-opacity-50 border border-l-indigo-50 items-center justify-center">
                    <h1 className="text-white font-bold text-4xl text-justify p-10">Controle e Utilize Estacionamentos de Forma Rápida e Simples</h1>
                    <div className="container h-auto w-auto justify-center px-10">
                        <button className="text-white bg-blue-900 ">
                            <span className=" h-auto w-auto p-3">Comece Aqui</span>
                        </button>
                    </div>
                </div>
            </Container>
        </section>
    )
}