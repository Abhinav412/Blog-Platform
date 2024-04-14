// main.js
document.getElementById('commentForm').addEventListener('submit', function(event) {
    event.preventDefault();
    var content = document.getElementById('commentContent').value;
    fetch('http://localhost:8080/comments', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({content: content}),
    })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch((error) => console.error('Error:', error));
  });
  
  fetch('http://localhost:8081/comments')
  .then(response => response.json())
  .then(data => {
    var commentsDiv = document.getElementById('comments');
    data.forEach(function(comment) {
      var p = document.createElement('p');
      p.textContent = comment.content;
      commentsDiv.appendChild(p);
    });
  })
  .catch((error) => console.error('Error:', error));