import { Button } from "@/components/ui/button";
import { PersonIcon } from "@radix-ui/react-icons";
export function AccountButton() {
    return (
        <Button 
        type="button" 
        variant={"ghost"} 
        className="text-white font-extrabold" 
        size={'icon'} 
        > <PersonIcon className="size-4 sm:size-6"/>
        </Button>
    );
}