# Distribute CAR files

Now it's time to distribute CAR files to storage providers so they can import on their side. Start by running the content provider service and download any Pieces for the dataset we have prepared:

```sh
singularity run content-provider
wget 127.0.0.1:8088/piece/bagaxxxx
```

If you have previously specified an output directory for exporting the CAR (which disables inline preparation), the CAR file will be served directly from those CAR files. Otherwise, if you have been using inline preparation or have accidentally deleted those CAR files, it will be served from the original data source directly
