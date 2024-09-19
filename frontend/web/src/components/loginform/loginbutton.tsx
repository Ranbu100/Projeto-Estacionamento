import Link from "next/link";
import { Button } from "../ui/button";

export function Loginbutton(){
    return (
        <Button variant="outline" className="text-blue-900" asChild>
            <Link href="/">Login</Link>
        </Button>
    )
}