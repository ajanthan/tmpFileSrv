// jshint devel:true
'use strict';

var source = $('#list-template').html();
/*eslint-disable no-undef*/
Handlebars.registerHelper('formatDate', function(dateStr) {
    // This guard is needed to support Blog Posts without date
    // the takeway point is that custom helpers parameters must be present on the context used to render the templates
    // or JS error will be launched
    var displayStr = '';
    var date = new Date(dateStr);

    if (typeof (date) === 'undefined') {
        return 'Unknown';
    }
    var now = new Date();
    var diff = now - date;
    diff = diff / 1000;
    if (diff < 60) {
        displayStr = displayStr + Math.round(diff) + ' seconds ago';
    } else {
        diff = diff / 60;
        if (diff < 60) {
            displayStr = displayStr + Math.round(diff) + ' minutes ago';
        } else {
            diff = diff / 60;
            if (diff < 24) {
                displayStr = displayStr + Math.round(diff) + ' hours ago';
            } else {
                diff = diff / 24;
                displayStr = displayStr + Math.round(diff) + ' days ago';
            }
        }
    }
    // These methods need to return a String
    return displayStr;
});
/*eslint-disable no-undef*/
var template = Handlebars.compile(source);

function listFiles() {
    $.ajax({
        url: '../files/',
        success: function(filesArray) {
            $('#list-files').html('');
            filesArray = JSON.parse(filesArray);
            if (filesArray.length > 0) {
                //var context = {title: "My New Post", body: "This is my first post!"};
                $('#list-files').show();
                var files = {
                    'files': filesArray
                };
                var html = template(files);
                $('#list-files').html(html);
            }
        }
    });
}
/*eslint-disable no-unused-vars*/
function deleteFile(fileName) {
    $.ajax({
        url: '../files/' + fileName,
        type: 'DELETE',
        success: function() {
            listFiles();
        }
    });
}


$(document).ready(function() {

    var progress = $('.progress');
    var bar = $('.progress-bar');
    var percent = $('.percent');
    var status = $('#status');
    progress.hide();
    status.hide();

    $('#add_file').ajaxForm({
        beforeSend: function() {
            status.hide();
            progress.show();
            status.empty();
            var percentVal = '0%';
            bar.width(percentVal);
            percent.html(percentVal);
        },
        uploadProgress: function(event, position, total, percentComplete) {
            var percentVal = percentComplete + '%';
            bar.width(percentVal);
            percent.html(percentVal);
        },
        success: function() {
            var percentVal = '100%';
            bar.width(percentVal);
            percent.html(percentVal);
            status.show();
            status.html('Successfully uploaded');
        },
        complete: function() {
            var percentVal = '0%';
            bar.width(percentVal);
            percent.html(percentVal);
            progress.hide();
            listFiles();
        },
        error: function(xhr) {
            status.show();
            status.html(xhr.responseText);
        }
    });
    $('#upload_file').change(function() {
        //submit the form here
        $('#add_file').submit();
    });
    listFiles();
});
