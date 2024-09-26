import {
    AlertDialog,
    AlertDialogContent,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogAction
  } from "@/components/ui/alert-dialog";
  import { useRouter } from "next/navigation";
  
  interface AlertLoginConfirmProps {
    show: boolean;
    onClose(open: boolean): void;
  }
  
  export function AlertLoginConfirm({ show, onClose }: AlertLoginConfirmProps) {
    const router = useRouter();
    return (
      <div className="h-dvh w-dvw items-center justify-center">
        <AlertDialog open={show} onOpenChange={onClose}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Login Realizado com Sucesso</AlertDialogTitle>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogAction onClick={() => { router.push("/home") }} >Entrar</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
  
    );
  }
  