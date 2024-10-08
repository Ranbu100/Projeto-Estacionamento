Folha de Requisitos: Sistema de Estacionamento

1. Requisitos Funcionais
 
1. Cadastro e Autenticação de Usuários
   - O sistema deve permitir o cadastro de usuários (administradores e clientes). --OK
   - O sistema deve permitir o login e logout dos usuários. --OK
   ** - O sistema deve suportar a recuperação de senha. --OK

2. Gerenciamento de Vagas
   - O administrador deve ser capaz de cadastrar e gerenciar vagas de estacionamento (número da vaga, tipo de vaga: carro, moto, etc.). -- OK
   - O sistema deve exibir o status das vagas (disponível, ocupada, reservada). 
   - Deve ser possível bloquear/desbloquear vagas específicas (por exemplo, manutenção).

3. Controle de Entrada e Saída de Veículos
   - O sistema deve permitir o registro da entrada e saída dos veículos, incluindo horário, placa, e vaga ocupada.
   - O sistema deve calcular automaticamente o tempo de permanência do veículo no estacionamento.

4. Cobrança e Pagamento
   - O sistema deve permitir a configuração de tarifas (por hora, diária, mensal).
   - O sistema deve calcular o valor a ser pago com base no tempo de permanência do veículo.
   - O sistema deve oferecer diferentes formas de pagamento (dinheiro, cartão de crédito, débito, PIX).
   - Deve ser possível gerar um recibo ou comprovante de pagamento.

5. Reserva de Vagas
   - O sistema deve permitir que os clientes façam reservas de vagas com antecedência.
   - O administrador deve poder definir um tempo limite para reservas.
   - O cliente deve poder cancelar ou modificar uma reserva antes do horário agendado.

6. Relatórios
   - O sistema deve gerar relatórios de ocupação do estacionamento (histórico de entradas e saídas, vagas mais usadas).
   - O sistema deve gerar relatórios financeiros (receitas diárias, mensais, por tipo de pagamento).
   - O sistema deve permitir a exportação dos relatórios em formatos como PDF.

7. Notificações e Alertas
   - O sistema deve notificar o cliente e o administrador em caso de ocupação máxima ou outras emergências.

8. Suporte a Dispositivos Móveis
   - O sistema deve ser responsivo, permitindo o uso em dispositivos móveis.
   - Deve haver a opção de desenvolver um aplicativo mobile com funcionalidades básicas (reserva de vagas, verificação de disponibilidade).

9. Monitoramento em Tempo Real
   - O sistema deve permitir ao administrador monitorar a ocupação em tempo real.


2. Requisitos Não Funcionais

1. Segurança
   - O sistema deve garantir que os dados dos usuários sejam armazenados de forma segura (criptografia de senhas e informações sensíveis). --OK
   - O sistema deve ter mecanismos de controle de acesso para diferentes níveis de usuários (administradores, operadores, clientes).

2. Desempenho
   - O sistema deve ser capaz de processar múltiplas entradas e saídas simultaneamente, suportando um grande número de vagas.
   - O tempo de resposta para operações deve ser inferior a 3 segundos.

3. Escalabilidade
   - O sistema deve ser capaz de escalar horizontalmente, permitindo a adição de novos módulos ou aumento de vagas sem perda de desempenho.

4. Usabilidade
   - O sistema deve ser intuitivo e fácil de usar, com uma interface clara para clientes e administradores. --OK

5. Compatibilidade
   - O sistema deve ser compatível com os principais navegadores (Chrome, Firefox, Edge) e com diferentes tamanhos de tela. --OK

6. Manutenibilidade
   - O código do sistema deve seguir padrões de desenvolvimento que facilitem a manutenção e evolução do sistema.
   - Deve haver documentação técnica disponível para desenvolvedores e administradores.

   - Deve exibir um mapa visual do estacionamento, indicando vagas ocupadas e disponíveis.
