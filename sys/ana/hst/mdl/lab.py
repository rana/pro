from __future__ import absolute_import, division, print_function

import os
import time
import numpy as np
import pandas as pd
import tensorflow as tf
from tensorflow import keras
import matplotlib.pyplot as plt

print(__file__ + "\n")

def load_data(path='data.npz', test_split=0.2, seed=113):
  assert 0 <= test_split < 1
  print("path:" + path + "\n")
  if path == 'data.npz':
    filepath = os.getcwd() +"/" + path
  else:
    filepath = path
  print("filepath:" + filepath + "\n")
  with np.load(filepath) as f:
    lbls = f['lbls']
    ftrs = f['ftrs']
    ftrNames = f['ftrNames']
  assert len(lbls) == len(ftrs)

  lbls_train = np.array(lbls[:int(len(lbls) * (1 - test_split))])
  ftrs_train = np.array(ftrs[:int(len(lbls) * (1 - test_split))])
  lbls_test = np.array(lbls[int(len(lbls) * (1 - test_split)):])
  ftrs_test = np.array(ftrs[int(len(lbls) * (1 - test_split)):])

  # shuffle train data only; emulates trading on later times
  np.random.seed(seed)
  indices = np.arange(len(lbls_train))
  np.random.shuffle(indices)
  lbls_train = lbls_train[indices]
  ftrs_train = ftrs_train[indices]

  return (lbls_train, ftrs_train), (lbls_test, ftrs_test), ftrNames

(lbls_train, ftrs_train), (lbls_test, ftrs_test), ftr_names = load_data('/home/rana/data/ml/ftrs.npz')


start = time.time()

# Shuffle the training set
order = np.argsort(np.random.random(lbls_train.shape))
ftrs_train = ftrs_train[order]
lbls_train = lbls_train[order]

print("Fiting set: {}".format(ftrs_train.shape))
print("Testing set:  {}".format(ftrs_test.shape))
print(ftrs_train[0])  # Display sample features, notice the different scales




df = pd.DataFrame(ftrs_train, columns=ftr_names)
df.head()

print(lbls_train[0:10])  # Display first 10 entries
print(ftrs_train[0])  # First training sample, normalized

node_cnt = int(len(ftrs_train)/len(ftr_names)+1) * 2
# node_cnt=64
print(">>> node_cnt {}".format(node_cnt))
def build_model():
  model = keras.Sequential([
    keras.layers.Dense(node_cnt, activation=tf.nn.relu,
                       input_shape=(ftrs_train.shape[1],)),
    keras.layers.Dense(node_cnt/2, activation=tf.nn.relu),
    # keras.layers.Dense(node_cnt/8, activation=tf.nn.relu),
    # keras.layers.Dense(node_cnt/4, activation=tf.nn.relu),
    # keras.layers.Dense(node_cnt/8, activation=tf.nn.relu),
    keras.layers.Dense(1)
  ])
  optimizer = tf.train.RMSPropOptimizer(0.001)
  # optimizer = tf.train.GradientDescentOptimizer(0.01)
  model.compile(loss='mse',
                optimizer=optimizer,
                metrics=['mae'])
  return model

model = build_model()
model.summary()

# Display training progress by printing a single dot for each completed epoch
class PrintDot(keras.callbacks.Callback):
  def on_epoch_end(self, epoch, logs):
    if epoch % 100 == 0: print('')
    print('.', end='')

EPOCHS = 2000

# The patience parameter is the amount of epochs to check for improvement
early_stop = keras.callbacks.EarlyStopping(monitor='val_loss', patience=20)
# Store training stats
history = model.fit(ftrs_train, lbls_train, epochs=EPOCHS,
                    validation_split=0.2, verbose=0,
                    callbacks=[ PrintDot()])



def plot_history(history):
  plt.figure()
  plt.xlabel('Epoch')
  plt.ylabel('Mean Abs Error')
  plt.plot(history.epoch, np.array(history.history['mean_absolute_error']),
           label='Fit Loss')
  plt.plot(history.epoch, np.array(history.history['val_mean_absolute_error']),
           label = 'Val loss')
  plt.legend()
  plt.ylim([0, 5])
plot_history(history)

[loss, mae] = model.evaluate(ftrs_test, lbls_test, verbose=0)

print("\n") 
print("Testing set Mean Abs Error: {:7.2f}".format(mae))
print("Testing set loss: {:7.2f}".format(loss))
print("ellapsed: {:.0f}s".format(time.time() - start))
