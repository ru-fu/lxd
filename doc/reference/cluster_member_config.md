(cluster-member-config)=
# Cluster member configuration

Each cluster member has its own key/value configuration with the following supported namespaces:

- `user` (free form key/value for user metadata)
- `scheduler` (options related to how the member is automatically targeted by the cluster)

The following keys are currently supported:

% Include content from [../config_options.txt](../config_options.txt)
```{include} ../config_options.txt
    :start-after: <!-- config group cluster start -->
    :end-before: <!-- config group cluster end -->
```

## Related topics

{{clustering_how}}

{{clustering_exp}}
