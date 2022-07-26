
$('#formulario-cadastro').on("submit", criarUsuario)

function criarUsuario(event) {
  event.preventDefault()

  if ($('#senha').val() != $('#confirmsenha').val()) {
    alert("As senha não coincidem!")
    return
  }

  $.ajax({
    url: "/usuarios",
    method: "post",
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $("#senha").val()
    }
  }).done (() => {
    console.log("criado usuario")
  }).fail(() => {
    console.error("erro ao criar o usuario")
  })

}