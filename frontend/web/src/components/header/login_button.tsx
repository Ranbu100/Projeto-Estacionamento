'use client'
import Link from "next/link"
import { Button } from "../ui/button"
export function Login(){

    return (
        <Button
        id="logar"
        className="flex items-center font-semibold text-blue-900 h-auto w-auto p-0 lg:p-2 bg-white border rounded-sm"
        asChild
        >
            <Link href={"/login"} className="h-auto w-auto p-0.5">Logar</Link>
        </Button>
    )
}