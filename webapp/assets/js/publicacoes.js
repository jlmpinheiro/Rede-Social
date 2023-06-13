// @ts-nocheck
$('#nova-publicacao').on('submit', criarPublicacao); //# para uma ID

//$('.curtir-publicacao').on('click', curtirPublicacao); //. para uma classe
$(document).on('click', '.curtir-publicacao', curtirPublicacao)
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao)
$(document).on('click', '.deletar-publicacao', deletarPublicacao)

$('#atualizar-publicacao').on('click', atualizarPublicacao); //# para uma ID



function criarPublicacao(evento) {
    evento.preventDefault(); 
    $.ajax({
        url: "/publicacao",
        method: "POST",
        data: {
            titulo:$('#titulo').val(),
            conteudo:$('#conteudo').val(),
        }
    }).done(function(){ 
        window.location="/home";
    }).fail(function(erro) {
        console.log(erro)
        swal.fire('Erro','Erro tentar publicar!','error')
    }); 
}

function curtirPublicacao(evento) {
    evento.preventDefault(); 

    const elementoClicado=$(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop("disabled",true);
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(){ 
        const contadorDeCurtidas=elementoClicado.next('span');
        const quantidadeDeCurtidas=parseInt(contadorDeCurtidas.text());
        
        contadorDeCurtidas.text(quantidadeDeCurtidas + 1 );

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');

    }).fail(function(erro) {
        console.log(erro);
        swal.fire('Erro','Erro tentar curtir!','error')
    }).always(function(){
        elementoClicado.prop("disabled",false);
    }); 
}

function descurtirPublicacao(evento) {
    evento.preventDefault(); 

    const elementoClicado=$(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop("disabled",true);
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function(){ 
        const contadorDeCurtidas=elementoClicado.next('span');
        const quantidadeDeCurtidas=parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1 );

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger')
        elementoClicado.addClass('curtir-publicacao');

    }).fail(function(erro) {
        console.log(erro);
        swal.fire('Erro','Erro tentar descurtir!','error')
    }).always(function(){
        elementoClicado.prop("disabled",false);
    }); 
}

function atualizarPublicacao() {
    $(this).prop('disabled', true)
    
    const publicacaoId=$(this).data('publicacao-id')
    console.log(publicacaoId)

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo:$('#titulo').val(),
            conteudo:$('#conteudo').val(),
        }
    }).done(function(){ 
        swal.fire('Sucesso!','Publicação atualizada com sucesso!','success').then(function() {
            window.location="/home";
        });
    }).fail(function(erro) {
        console.log(erro)
        swal.fire('Erro','Erro tentar publicar!','error')
    }).always(function(){
        $('#atualizar-publicacao').prop('disabled', false)
    }); 
}

function deletarPublicacao(evento) {
    evento.preventDefault(); 

    swal.fire({
        title: 'Atenção',
        text:'Tem certeza que deseja excluir essa Publicação? Essa ação é irreversível!',
        showCancelButton: true,
        cancelButtonText:'Cancelar',
        icon: 'warning'
        }).then(function(confirmacao) {
            if (!confirmacao.value) return

            const elementoClicado=$(evento.target);
            const publicacao = elementoClicado.closest('div')
            const publicacaoId = publicacao.data('publicacao-id');
            elementoClicado.prop("disabled",true);
            
            $.ajax({
                url: `/publicacoes/${publicacaoId}`,
                method: "DELETE"
            }).done(function(){ 
                publicacao.fadeOut("slow",function(){
                    $(this).remove();
                });
            }).fail(function(erro) {
                swal.fire('Erro','Erro tentar excluir!','error')
                elementoClicado.prop("disabled",false);
            }).always(function(){
                elementoClicado.prop("disabled",false);
            });
    });

 
}
