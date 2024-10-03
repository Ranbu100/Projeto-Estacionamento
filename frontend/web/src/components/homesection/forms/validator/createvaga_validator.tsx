import * as yup from "yup";

export const formSchemaCreate = yup.object().shape({
    numero_vaga: yup.number().required("Coloque o número da vaga"),
    tipo_vaga: yup.string().required("É obrigatório colocar o tipo da vaga"),
});
