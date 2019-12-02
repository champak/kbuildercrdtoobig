# kbuildercrdtoobig
Illustrates an error from kubebuilder
kubebuilder (v <= 2.0.0) appears to include the entire spec for native types such as Pods into the generated CRD. This results 
in issues with using "kubectl apply" to install the CRD as implemented by the make install step; kubectl uses annotations to
store the CRD and we blow past the max annotation limit. This repo is a simple illustration of the issue.
Discussions are on the kubebuilder issues list (https://github.com/kubernetes-sigs/kubebuilder/issues/1140)
