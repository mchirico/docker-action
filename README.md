# docker-action


```yaml
name: Docker Action
on: [push]

jobs:
  hello_world_job:
    runs-on: ubuntu-latest
    name: A job to say hello
    steps:
    - name: Hello world action step
      id: hello
      uses: mchirico/docker-action@v1
      with:
        who-to-greet: 'Mona the Octocat'
    # Use the output from the `hello` step
    - name: Get the output time
      run: echo "The time was ${{ steps.hello.outputs.time }}"

```
# Sample output:

![image](https://user-images.githubusercontent.com/755710/75095859-453f7300-5567-11ea-8203-9169a9ad45fb.png)
