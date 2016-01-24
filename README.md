# **Servidor**

Servidor is a light-weight no-database git server following git smart HTTP protocol. You can do all kind of git remote operations like push, pull, fetch and clone. Host the server very easily and get started.

## Features supported as of now

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


  ![](https://github.com/gophergala2016/servidor/blob/master/screencasts/git_operations_without_auth.gif)

  [More screencasts](https://github.com/gophergala2016/servidor/tree/master/screencasts)

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

Start the server by executing `servidor` binary. By default, servidor will listen to http://localhost:8000 for incoming requests.


## Options:
```
./servidor -h
Usage of ./servidor:
  -R	Set Whether ReceivePack(push operation) will be restricted
  -U	Set Whether UploadPack(clone, pull, fetch operations) will be restricted
  -a	Enable basic authentication for all http operations
  -b string
    	Hostname to be used (default "0.0.0.0")
  -c string
    	Set the path from where the password file is to be read(to be set whenever -a flag is used)
  -g string
    	Mention the gitPath if its different on hosting machine (default "/usr/bin/git")
  -p string
    	Port on which servidor will listen (default "8000")
  -r string
    	Set the path where repositories will be saved, Just mention the base path("repos" directory will be automatically created inside it) (default "/home/administrator/servidor")
  -s	Enable tls connection
```

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

## Setup for extra features

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

- To enable ssl connection

 *Generated private key*
  ```
  $ openssl genrsa -out server.key 2048
  ```

  *Generate the certificate*
  ```
  $ openssl req -new -x509 -key server.key -out server.pem -days 3650
  ```

 Since the certificates are self authorized, server verification must be turned off for clients:
 
 For curl, `use -k flag`

 For git, `export GIT_SSL_NO_VERIFY=1`

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

## Feature to come

- Webhook support
- More repo metrics

## License

MIT, see the LICENSE file.
