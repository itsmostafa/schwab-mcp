from datetime import datetime
import mongoengine as mongo


def global_init():
    mongo.register_connection(name="analytics", alias="core")