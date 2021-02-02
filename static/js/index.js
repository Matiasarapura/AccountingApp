


$(document).ready(function(){
    getBalance()
    fillTransactions()
    // click on button submit
    $("#submit").on('click', function(){
        // send ajax
        var data = {
            Amount : parseFloat($("#Amount").val()),
            Type: $("#Type").val()
        }
        $.ajax({
            url: 'http://localhost:8080/transaction',
            type : "POST",
            dataType : 'json',
            data :JSON.stringify(data) ,
            success : function(result) {

                console.log(result);
                getBalance()
                fillTransactions()
            },
            error: function(xhr, resp, text) {
                console.log(xhr, resp, text);
                alert("Something went wrong")
            }
        })
    });
});

function getBalance(){
    $.getJSON( "http://localhost:8080/balance", function( data ) {
        console.log(data)
        $("#balance").html(data.balance);
    })
}

function fillTransactions(){
    html = ''
    $.getJSON( "http://localhost:8080/transaction", function( data ) {
        console.log(data)
        for(var i = 0; i < data.length; i++)
            html += '<tr><td>' + data[i].Id + '</td><td>' + data[i].Date + '</td><td>' + data[i].Amount + '</td><td>' + data[i].Type + '</td></tr>';
        $('#transactions tr').first().after(html);
    })
}


