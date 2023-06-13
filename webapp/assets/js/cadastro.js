// @ts-nocheck
$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault(); 

    //alert("senha:"+$('#senha').val() +" senha1= "+$('#confirmar-senha').val())
    if ($('#senha').val() != $('#confirmar-senha').val()) {
        swal.fire('Ops...!','As senhas não coincidem!!','error')
            return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome:$('#nome').val(),
            nick:$('#nick').val(),
            email:$('#email').val(),
            senha:$('#senha').val(),
        }
    }).done(function(){ 
        swal.fire('Sucesso!','Usuário cadastrado com sucesso!!!','success').then(function() {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email:$('#email').val(),
                    senha:$('#senha').val(),
                }
            }).done(function(){ 
                window.location="/home";
            }).fail(function(erro) {
                swal.fire('Erro','Erro ao tentar logar!','error')
            });
        })
    }).fail(function(erro) {
        swal.fire('Erro','Erro ao cadastrar usuário!','error')
    }); 
}