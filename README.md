# **Servidor**

Servidor is a light-weight no-database git server following git smart HTTP protocol. You can do all kind of git remote operations like push, pull, fetch and clone. Host the server very easily and get started. Features supported as of now are listed below:

- Git Remote Operations
  - [x] Cloning of git repository.
  - [x] Push operation.
  - [x] Fetch operation.
  - [x] Pull operation.
- APIs
  - [x] Create git repository - POST
  - [x] List all repositories corresponding to a user. - GET
  - [x] List a particular repository corresponding to a user. -GET
  - [x] List all the branches in a repository. - GET
  - [x] List a particular branch in a repository. - GET
- Extra Features
  - [x] Basic authentication as per flag
  - [x] Allowing TLS connection as per flag.
  - [x] Restricting push, pull operations as per flag.

## Motivation

While setting up our project a few months back, we had to go through the trouble of setup and configuration
needed in GitLab. To do away with all of that in future, I decided to create a git server of my own. It will typically help small group of coders, who wish to maintain private repositories within a local network and don't want to do all sorts of setup needed in GitLab and other providers.

## Installation

- Install cmake from brew(Mac) or apt-get(linux) package manager. To build from source, follow [this](https://cmake.org/install/) link

- Install libgit2 as follows :
    ```
      $ wget https://github.com/libgit2/libgit2/archive/v0.23.4.tar.gz
      $ tar xzf v0.23.4.tar.gz
      $ cd libgit2-0.23.4/
      $ cmake .
      $ make
      $ sudo make install
    ```

- Build the project

  Assuming you have installed a recent version of
  [Go](https://golang.org/doc/install), you can simply run

  ```
  go get github.com/gophergala2016/servidor
  ```

  This will download Servidor to `$GOPATH/src/github.com/gophergala2016/servidor`. From
  this directory run `go build` to create the `servidor` binary.

- Troubleshooting:-  
    ```ImportError: libgit2.so.0: cannot open shared object file: No such file or directory```  
     This happens for instance in Ubuntu, the libgit2 library is installed within the `/usr/local/lib` directory, but the linker does not look for it there.
     To fix this call
    ```
      $ sudo ldconfig
    ```

## Getting started

Start the server

```
servidor  [-a] [-b] [-c] [-g] [-h] [-p] [-r] [-s] [-R] [-U]
```

## Options:

``` -a ```
    Enables basic authentication. If enabled, clients must provide valid username and password to perform git operations and create repositories.

``` -b host-name ```
    Sets hostname for Servidor.

``` -c path/to/password/file ```
     Used only if **-a** flag is set. It is used to specify the location of password file where list of authorized users and password is maintained.
     To generate the password, commandline utility tool [htpasswd](https://httpd.apache.org/docs/2.2/programs/htpasswd.html) has been used.
     Examples described later, show elaborate usage of htpasswd to create SHA-1 hash of the password.

``` -g /path/to/git ```
    Sets the git path. Default is "/usr/bin/git".

``` -h ```
    Usage of flags.

``` -r path/to/save/repos ```
     Sets the repository path where the repositories of various users will be saved. If not specified, present working directory will be set as the default path for saving repositories.

``` -p port-number ```
     Sets the port on which Servidor will listen.

``` -s ```
     Enables ssl connection.

``` -R ```
     Restricts ReceivePack(push operation)

``` -U ```
     Restricts UploadPack(clone, pull, fetch operations)

## Setup for extra features

- To enable ssl connection

  *Generated private key*
   ```
   $ openssl genrsa -out server.key 2048
   ```

   *Generate the certificate*
   ```
   $ openssl req -new -x509 -key server.key -out server.pem -days 3650
   ```
- To enable basic authentication, create the password file as follows

   *The password file has entries of the format : ```username:password(in SHA-1 encoded format)```*

   - To generate the password file, use htpasswd tool. Install it by using
    ```
    $ sudo apt-get install apache2-utils
    ```

   - Once installed you can use
    ```
     $ htpasswd -cs path/to/create/password/file/filename username1

     $ htpasswd -s path/to/create/password/file/filename username2
    ```

   Note: while creating sha-1 password for the second user, do not use -c flag. It is used to create the file for the first time. See [documentation](https://httpd.apache.org/docs/2.2/programs/htpasswd.html) if needed.

## Usages:

- Create git repository as follows :

  ```
  $ curl -X POST http://<hostname>:<port>/api/repos/create
  -d '{"username":"username1","reponame":"project1"}'
  ```

    - Typical successful response :
    ```
    {
      "response_message": "repository created successfully",
      "clone_url": "http://<hostname>:<port>/username1/project1.git"
    }
    ```

    - Typical unsuccessful response :
    ```
    {
       "response_message": "repository already exists for user",
       "clone_url": "http://<hostname>:<port>/username1/project1.git"
    }
    ```

- Now, Clone the repository using the clone_url. Do stuffs, push changes set to remote, pull changes from remote etc.

## API References

- To display the list of APIs available
  ```
  $ curl http://<hostname>:<port>
  ```
  *Response:*

  ```
  {
    "create_repo_url": "http://<hostname>:<port>/api/repos/create",
    "user_repositories_url": "http://<hostname>:<port>/api/{user-name}/repos",
    "user_repository_url": "http://<hostname>:<port>/api/{user-name}/repos/{repo-name}",
    "branches_url": "http://<hostname>:<port>/api/{user-name}/repos/{repo-name}/branches{/branch-name}"
  }
  ```
