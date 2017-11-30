# Python Project with Pipfile.lock

This project is an example of a python project that uses Pipfile.lock to manage its dependencies.

#### Install dependencies
This project requires `Pipenv` to install dependencies. In your virtual environment, run `pipenv install --ignore-pipfile`. This command will install dependencies from the lock file.

`Pipfile.lock` at the root directory declares the dependencies:
- django==1.7.1
- feedparser==5.1.1
- pyasn1==0.4.2
- pycrypto==2.4
- pyjwt==0.4.2
- raven==1.9.4
- requests==2.2.1
- rsa==3.4
- simplejson==2.4.0

Total number of dependencies:
- 9 direct dependencies
- 28 vulnerabilities

#### Vulnerable Call Chain
`main.py` has a method call that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `with-pipfile-dot-lock` project, you should get results that are coherent with the details above if the tool supports using Pipfile.lock to manage python dependencies.