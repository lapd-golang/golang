# Exploration

### Existing Solutions

1. [kontinuous](https://github.com/AcalephStorage/kontinuous)
   
    - define the pipeline kind resource and specification;
    - use the job kind resource to do the work
    - various agents for the detail work
    - REST API with the clients for key/value from background.
    - The content of deployment:
        * use the kubernetes namespace,deployment and service resource on the basis of image build
    - The pipeline build:
        * how to trigger the build?
    - Use the kubernetes job to do the command task;
    - Leverage the key/value storage and contruct necessary data schema and models

    
2. Ansible Integration

    * the main point is to use ansible module to manage kubernetes's resources
  
3. [AppController](https://github.com/Mirantis/k8s-AppController) 

    - reshape the model on the base of existing models
    - use config map to define the dependencies and create the dependent graph
    - run the dependent graph
  
4. [cycle-clone](https://github.com/caicloud/cyclone)

  - use cmd and api for data acceptance;
  - store for data persitence;
  - workers for task execution;
  - event for data insertion and update;
    
4. to-do-list

  - [good article](https://github.com/debalin/devops-kubernetes)
  - [deep introduction](https://www.qstack.com/devops-12-factor-application/)
