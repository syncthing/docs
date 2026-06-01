Package *build*
~~~~~~~~~~~~~~~

Metric *syncthing_build_info* (gauge vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

A metric with a constant '1' value labeled by version information from
when the binary was built.

Package *connections*
~~~~~~~~~~~~~~~~~~~~~

Metric *syncthing_connections_active* (gauge vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Number of currently active connections, per device. If value is 0, the
device is disconnected.

Package *db*
~~~~~~~~~~~~

Metric *syncthing_db_files_updated_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of files updated.

Metric *syncthing_db_operation_seconds_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total time spent in database operations, per folder and operation.

Metric *syncthing_db_operations_current* (gauge vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Number of database operations currently ongoing, per folder and
operation.

Metric *syncthing_db_operations_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of database operations, per folder and operation.

Package *events*
~~~~~~~~~~~~~~~~

Metric *syncthing_events_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of created/forwarded/dropped events.

Package *fs*
~~~~~~~~~~~~

Metric *syncthing_fs_operation_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of filesystem bytes transferred, per filesystem root and
operation.

Metric *syncthing_fs_operation_seconds_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total time spent in filesystem operations, per filesystem root and
operation.

Metric *syncthing_fs_operations_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of filesystem operations, per filesystem root and
operation.

Package *model*
~~~~~~~~~~~~~~~

Metric *syncthing_model_folder_conflicts_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of conflicts.

Metric *syncthing_model_folder_processed_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total amount of data processed during folder syncing, per folder ID and
data source (network/local_origin/local_other/skipped).

Metric *syncthing_model_folder_pull_seconds_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total time spent in folder pull iterations, per folder ID.

Metric *syncthing_model_folder_pulls_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of folder pull iterations, per folder ID.

Metric *syncthing_model_folder_scan_seconds_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total time spent in folder scan iterations, per folder ID.

Metric *syncthing_model_folder_scans_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of folder scan iterations, per folder ID.

Metric *syncthing_model_folder_state* (gauge vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Current folder state.

Metric *syncthing_model_folder_summary* (gauge vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Current folder summary data (counts for global/local/need
files/directories/symlinks/deleted/bytes).

Package *protocol*
~~~~~~~~~~~~~~~~~~

Metric *syncthing_protocol_recv_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total amount of data received, per device.

Metric *syncthing_protocol_recv_decompressed_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total amount of data received, after decompression, per device.

Metric *syncthing_protocol_recv_messages_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of messages received, per device.

Metric *syncthing_protocol_sent_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total amount of data sent, per device.

Metric *syncthing_protocol_sent_messages_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of messages sent, per device.

Metric *syncthing_protocol_sent_uncompressed_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total amount of data sent, before compression, per device.

Package *scanner*
~~~~~~~~~~~~~~~~~

Metric *syncthing_scanner_hashed_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total amount of data hashed, per folder.

Metric *syncthing_scanner_scanned_items_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total number of items (files/directories) inspected, per folder.

