# **Servidor**

Servidor is a light-weight no-database git server following git smart http protocol. Host the server hastle free and get started. It will typically help small group of coders, who wish to maintain private repositories within a local network and don't want to do all kind of setup and configuration needed in GitLab and other providers.

More features will be added in coming hours and also in days to come, to achieve completion.
I will be adding more documentation on regular basis.

## Installation

- Install cmake from brew or apt-get package manager. To build from source, follow [this](https://cmake.org/install/) link

- Install libgit2 as follows :
    ```
      $ wget https://github.com/libgit2/libgit2/archive/v0.23.4.tar.gz
      $ tar xzf v0.23.4.tar.gz
      $ cd libgit2-0.23.4/
      $ cmake .
      $ make
      $ sudo make install
    ```

- Install go packages as folloes :
    ```
      $ go get github.com/libgit2/git2go
      $ go get github.com/gorilla/mux
    ```

- Build the project using :
    ```
      $ go build
    ```

- Troubleshooting:-  
    ```ImportError: libgit2.so.0: cannot open shared object file: No such file or directory```  
         This happens for instance in Ubuntu, the libgit2 library is installed within the /usr/local/lib directory, but the linker does not look for it there.
         To fix this call
    ```
      $ sudo ldconfig
    ```

## Usage

Start the server.

```
servidor -h will show the flags to start the server
```

- Create git repository as follows :
**$ curl -X POST `http://<hostname>:<port>/api/repos/create` -d '{"username":"username1","reponame":"project1"}'**

    - Typical response : {"response_message":"Repository created successfully","clone_url":"`http://<hostname>:<port>/username1/project1.git`"}

- Now, Clone the repository using the clone_url. Do stuffs, push change sets to remote, pull changes from remote etc.

## API References
- **$ curl `http://<hostname>:<port>/`** - shows the list of APIs available

# TODO
- Git Operations
  - [x] Allow Cloning of git repository.
  - [x] Allow Push operation.
  - [x] Allow Fetch operation.
  - [x] Allow Pull operation.
- APIs
  - [x] Create git repository - POST
  - [x] List all repositories corresponding to a user. - GET
  - [x] List a particular repository corresponding to a user. -GET
  - [x] List all the branches in a repository. - GET
  - [x] List a particular branch in a repository. - GET
- Extra Features
  - [x] Provide basic authentication
  - [ ] Enable ssl connection as per configuration.
  - [ ] Block push, pull operations as per configuration.
