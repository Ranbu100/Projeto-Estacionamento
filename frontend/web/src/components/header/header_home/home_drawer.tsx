import { Button } from '@/components/ui/button';
import { useState } from 'react';
import { ListBulletIcon } from "@radix-ui/react-icons";
import { AccountButton } from './account_button';

export function HomeDrawer() {
    const [isOpen, setIsOpen] = useState(false);

    const toggleDrawer = () => {
        setIsOpen(!isOpen);
    };

    return (
        <div>
            <Button variant={"ghost"} onClick={toggleDrawer} className='text-white'>
                <ListBulletIcon className='size-6' />
            </Button>

            {isOpen && (
                <div
                    className="fixed inset-0 bg-black bg-opacity-50 z-40"
                    onClick={toggleDrawer}
                />
            )}

            <div
                className={`fixed top-0 left-0 h-full w-64 bg-white shadow-lg z-50 transform transition-transform duration-300 ${isOpen ? 'translate-x-0' : '-translate-x-full'}`}
            >
                <div className="flex flex-col h-full justify-between  p-4">
                   <div className='flex flex-col p-4 space-y-3'>
                   <AccountButton />
                   </div>
                    
                    
                    <Button onClick={toggleDrawer} variant={'ghost'} className=" text-red-500">
                        Fechar
                    </Button>
                </div>
            </div>
        </div>
    );
}
