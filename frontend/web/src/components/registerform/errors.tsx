import {
    AlertDialog,
    AlertDialogContent,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogAction
  } from "@/components/ui/alert-dialog";
  import { useRouter } from "next/navigation";
  
  interface AlertRegisterErrorsProps {
    show: boolean;
    onClose(open: boolean): void;
  }
  
  export function AlertRegisterErrors({ show, onClose }: AlertRegisterErrorsProps) {
    const router = useRouter();
    return (
      <div className="h-5/6 w-5/6 items-center justify-center bg-fixed">
        <AlertDialog open={show} onOpenChange={onClose}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Houve um erro no cadastro, tente novamente</AlertDialogTitle>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogAction onClick={() => { router.push("/register") }} >Tentar Novamente</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
  
    );
  }
  