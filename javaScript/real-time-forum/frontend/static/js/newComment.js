export const newCommentTemplate = () => {
   document.title = "New Comment",
    render(
         `<div class="container">
    <div class="header">
        <a href="/">
        <h1>Write your comment below</h1>
    </a>
        <form action="/">
            <button type="submit">Home</button>
        </form>
    </div>
    <br>
    <form action="/newComment" method="POST">
        <textarea type="text" name="content" required rows="4" cols="50" placeholder="Write your comment here"></textarea>
        <br>
    <input type="hidden" name="t" value="{{ . }}">
    <br>
   
    <input type="submit">
</form>
        </div>
    </div>
</div>`)
}