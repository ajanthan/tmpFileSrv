var source   = $("#list-template").html();
var template = Handlebars.compile(source);

$(document).ready(function (){
	$('#uploadstatus').hide()
	//$('#list-files').hide()
	//$('#percent').hide()
	//$('#bar').hide()
	//$('#message').hide()
	var options 	={
	beforeSend : function (){
		$('#uploadstatus').show()
		$('#progress').show()
		$('#bar').width('0%')
		$('#percent').html('0%')
		$('#message').val('Upload is started')
	},
	uploadProgress : function(event, position, total, percentComplete) {
		$('#bar').width(percentComplete+'%')
		$('#percent').html(percentComplete+'%')
	},
	success: function() {
		$('#bar').width('100%')
		$('#percent').html('100%')
	},
	complete: function(response) {
		  $("#message").html("<font color='green'>"+"sucess"+"</font>");
		listFiles()
	},
	error: function(){
		$("#message").html("<font color='red'> ERROR: unable to upload files</font>");
		listFiles()
	}	
	};
	$("#fileUploadForm").ajaxForm(options);
	listFiles()
});

function listFiles(){
	$.ajax(
		{
		url:"../files/",
		success : function (filesArray){
			$('#list-files').html('');
		filesArray=JSON.parse(filesArray)
		if(filesArray.length > 0){
			
		//var context = {title: "My New Post", body: "This is my first post!"};
		$('#list-files').show();
		var files={
			"files":filesArray
		};
         var html    = template(files);
		$('#list-files').html(html);
		
		}
		}	
		}
	);
	
}
function deleteFile(fileName){
$.ajax(
		{
		url:"../files/"+fileName,
		type:"DELETE",
		success : function (){
		listFiles()
		}}
	);	
}