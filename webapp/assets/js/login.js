// @ts-nocheck
$('#login').on('submit', FazerLogin);

function FazerLogin(evento) {
    evento.preventDefault(); 
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
}