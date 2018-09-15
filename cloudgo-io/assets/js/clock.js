$(document).ready(function() {
    setInterval(function() {
      var dateNtime = new Date();
      var year = dateNtime.getFullYear();
      var month = dateNtime.getMonth() + 1;
      var day = dateNtime.getDate();
      var hours = dateNtime.getHours();
      var minute = dateNtime.getMinutes();
      var second = dateNtime.getSeconds();

      var dateStr = month + "/" + day + ", " + year;
      var timeStr = hours + ":" + minute + ":" + second;
      $('.index-date').html("Today is " + dateStr);
      $('.index-time').html("The time is " + timeStr);
    }, 1000)
})
