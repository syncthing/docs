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

Metric *syncthing_model_folder_processed_bytes_total* (counter vector)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Total amount of data processed during folder syncing, per folder ID and
data source (network/local_origin/local_other/local_shifted/skipped).

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

