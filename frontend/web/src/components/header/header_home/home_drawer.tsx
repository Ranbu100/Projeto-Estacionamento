import { Button } from '@/components/ui/button';
import { useState } from 'react';
import { ListBulletIcon } from "@radix-ui/react-icons";

export function HomeDrawer() {
    const [isOpen, setIsOpen] = useState(false);

    const toggleDrawer = () => {
        setIsOpen(!isOpen);
    };

    return (
        <div>
            <Button variant={"ghost"} onClick={toggleDrawer} className='text-white'>
                <ListBulletIcon className='size-6'/>
            </Button>

            {isOpen && (
                <div
                    className="fixed inset-0 bg-black bg-opacity-50 z-40"
                    onClick={toggleDrawer}
                />
            )}

            {/* Drawer que desliza da direita */}
            <div
                className={`fixed top-0 left-0 h-full w-64 bg-white shadow-lg z-50 transform transition-transform duration-300 ${isOpen ? 'translate-x-0' : 'translate-x-full'
                    }`}
            >
                <div className="p-4">
                    <h2 className="text-lg font-semibold">Drawer Menu</h2>
                    <p>Conteúdo do Drawer</p>

                    {/* Botão para fechar o drawer */}
                    <Button onClick={toggleDrawer} className="mt-4">
                        Fechar Drawer
                    </Button>
                </div>
            </div>
        </div>
    );
}
